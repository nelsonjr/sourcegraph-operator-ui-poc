import {
  Button,
  FormControl,
  InputLabel,
  MenuItem,
  Paper,
  Select,
  Stack,
  Typography,
} from "@mui/material";
import { useState } from "react";
import { changeStage } from "./debugBar";

export const Install: React.FC = () => {
  const [version, setVersion] = useState<string>("5.3.1");

  const install = () => {
    changeStage({ action: "installing", data: version });
  };

  return (
    <div className="install">
      <Typography variant="h3">Install Sourcegraph Appliance</Typography>
      <Paper elevation={3} sx={{ p: 4 }}>
        <Stack direction="column" spacing={2} sx={{ alignItems: "center" }}>
          <FormControl sx={{ minWidth: 250 }}>
            <InputLabel id="demo-simple-select-label">Version</InputLabel>
            <Select
              value={version}
              label="Age"
              onChange={(e) => setVersion(e.target.value)}
            >
              <MenuItem value={"5.3.0"}>5.3.0</MenuItem>
              <MenuItem value={"5.3.1"}>5.3.1</MenuItem>
              <MenuItem value={"5.4.0 (beta)"}>5.4.0 (beta)</MenuItem>
            </Select>
          </FormControl>
          <Button variant="contained" sx={{ width: 200 }} onClick={install}>
            Install
          </Button>
        </Stack>
      </Paper>
    </div>
  );
};
