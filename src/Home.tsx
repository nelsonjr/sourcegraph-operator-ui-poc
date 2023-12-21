import { CircularProgress, Typography } from "@mui/material";
import "./App.css";
import { OutletContext } from "./Frame";
import { useOutletContext } from "react-router-dom";

export const Home: React.FC = () => {
  const context = useOutletContext<OutletContext>();

  return (
    <div className="home">
      <CircularProgress />
      {context.online || context.needsLogin ? (
        <>
          <Typography variant="h5">
            Operator connected. Please wait...
          </Typography>
        </>
      ) : (
        <>
          <Typography variant="h5">
            Please wait, while the Sourcegraph Appliance Operator connects...
          </Typography>
        </>
      )}
    </div>
  );
};
