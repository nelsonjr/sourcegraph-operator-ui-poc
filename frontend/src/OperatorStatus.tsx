import { Navigate } from "react-router-dom";
import { ContextProps } from "./Frame";

export const OperatorStatus: React.FC<ContextProps> = ({ context }) => {
  console.log(context);
  const Status = () =>
    context.online === undefined ? (
      <div className="status connecting">connecting</div>
    ) : context.online === true || context.needsLogin === true ? (
      <div className="status online">online</div>
    ) : (
      <div className="status offline">offline</div>
    );

  switch (context.stage) {
    case "refresh":
      document.location = "https://sourcegraph.com/search";
      break;
  }

  return (
    <div id="operator-status">
      Operator status: <Status />
      {context.online === false && <Navigate to="/" />}
      {context.stage === "unknown" && <Navigate to="/" />}
      {context.stage === "install" && <Navigate to="/install" />}
      {context.stage === "installing" && <Navigate to="/install/progress" />}
      {context.stage === "install-wait-admin" && (
        <Navigate to="/install/wait-for-admin" />
      )}
      {context.stage === "upgrading" && <Navigate to="/upgrade/progress" />}
      {context.stage === "maintenance" && <Navigate to="/maintenance" />}
    </div>
  );
};
