import { Typography } from "@mui/material";

import { ConcensusNavBar } from "./components/navbar";

function App() {
  return (
    <div className="App">
      <Typography variant="h1" component="p">
        <ConcensusNavBar />
      </Typography>
    </div>
  );
}

export default App;
