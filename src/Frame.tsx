import { Outlet } from "react-router-dom";
import logo from "../assets/sourcegraph-reverse-logo.png";
import { OperatorDebugBar } from "./OperatorDebugBar";
import { OperatorStatus } from "./OperatorStatus";
import { Typography } from "@mui/material";
import SailingIcon from "@mui/icons-material/Sailing";
import { useEffect, useState } from "react";

const FetchStateTimerMs = 1 * 1000;

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

interface StatusResult {
  online: boolean;
  stage?: stage;
}

const fetchStatus = async (): Promise<StatusResult> =>
  new Promise<StatusResult>((resolve) => {
    fetch("/api/operator/v1beta1/stage")
      .then((result) => {
        if (!result.ok) {
          resolve({ online: false });
          return;
        }
        return result;
      })
      .then((result) => result?.json())
      .then((result) => {
        resolve({ online: true, stage: result.stage });
      })
      .catch(() => {
        resolve({ online: false });
      });
  });

export interface OutletContext {
  online: boolean;
  stage: stage;
}

export const Frame: React.FC = () => {
  const [online, setOnline] = useState<boolean>(false);
  const [stage, setStage] = useState<stage>("unknown");

  useEffect(() => {
    const timer = setInterval(() => {
      fetchStatus().then((result) => {
        setOnline(result.online);
        if (result.online && result.stage) {
          setStage(result.stage);
        }
      });
    }, FetchStateTimerMs);
    return () => clearInterval(timer);
  }, []);

  const context: OutletContext = {
    online: online,
    stage: stage,
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
        <Outlet context={context} />
      </div>
      <OperatorDebugBar context={context} />
      <OperatorStatus context={context} />
    </div>
  );
};
