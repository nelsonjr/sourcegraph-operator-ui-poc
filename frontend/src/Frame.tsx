import SailingIcon from "@mui/icons-material/Sailing";
import { Typography } from "@mui/material";
import { useEffect, useState } from "react";
import { Outlet } from "react-router-dom";
import logo from "../assets/sourcegraph-reverse-logo.png";
import { Login } from "./Login";
import { OperatorDebugBar } from "./OperatorDebugBar";
import { OperatorStatus } from "./OperatorStatus";
import { adminPassword, call } from "./api";

const FetchStateTimerMs = 1 * 1000;
const WaitToLoginAfterConnectMs = 1 * 1000;

export type stage =
  | "unknown"
  | "install"
  | "installing"
  | "install-wait-admin"
  | "upgrading"
  | "maintenance"
  | "refresh";

export interface ContextProps {
  context: OutletContext;
}

export interface OutletContext {
  online: boolean;
  onlineDate?: number;
  stage?: stage;
  needsLogin?: boolean;
}

const fetchStatus = async (
  lastContext: OutletContext
): Promise<OutletContext> =>
  new Promise<OutletContext>((resolve) => {
    call("/api/operator/v1beta1/stage")
      .then((result) => {
        if (!result.ok) {
          if (result.status === 401) {
            resolve({
              online: false,
              needsLogin: true,
              onlineDate: lastContext.onlineDate ?? Date.now(),
            });
          } else {
            resolve({ online: false, onlineDate: undefined });
          }
          return;
        }
        return result;
      })
      .then((result) => result?.json())
      .then((result) => {
        resolve({
          online: true,
          stage: result.stage,
          onlineDate: lastContext.onlineDate ?? Date.now(),
        });
      })
      .catch(() => {
        resolve({ online: false, onlineDate: undefined });
      });
  });

export const Frame: React.FC = () => {
  const [context, setContext] = useState<OutletContext>({
    online: false,
  });
  const [login, setLogin] = useState<boolean>(false);
  const [password, setPassword] = useState<string>();
  const [failedLogin, setFailedLogin] = useState<boolean>(false);

  useEffect(() => {
    const timer = setInterval(() => {
      if (failedLogin) {
        setLogin(true);
      }

      fetchStatus(context).then((result) => {
        setContext(result);
        if (result.needsLogin) {
          setLogin(true);
          if (password !== undefined) {
            setFailedLogin(true);
          }
        } else {
          setLogin(false);
          setFailedLogin(false);
        }
      });
    }, FetchStateTimerMs);
    return () => clearInterval(timer);
  }, [password, failedLogin, context]);

  useEffect(() => {
    adminPassword.password = password;
  }, [password]);

  const doLogin = (p: string) => {
    setPassword(p);
    setFailedLogin(false);
  };

  return (
    <div id="frame">
      <header>
        <img id="logo" src={logo} />
        <div className="spacer" />
        <SailingIcon fontSize="large" />
        <Typography variant="h4">Nemo Ride Management</Typography>
      </header>
      <div id="content">
        {login &&
        context.onlineDate &&
        context.onlineDate < Date.now() - WaitToLoginAfterConnectMs ? (
          <Login onLogin={doLogin} failed={failedLogin} />
        ) : (
          <Outlet context={context} />
        )}
      </div>
      <OperatorDebugBar context={context} />
      <OperatorStatus context={context} />
    </div>
  );
};