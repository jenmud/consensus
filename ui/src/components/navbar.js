//import { makeStyles } from "@mui/material";
import { AppBar, Toolbar, Typography } from "@mui/material";
import React from "react";

//const useStyles = makeStyles((theme) => ({
//  root: {
//    flexGrow: 1,
//  },
//  menuButton: {
//    marginRight: theme.spacing(2),
//  },
//}));

export const ConcensusNavBar = () => {
  //const classes = useStyles();

  return (
    <AppBar position="static">
      <Toolbar variant="dense">
        <Typography variant="h6" color="inherit">
          Concensus
        </Typography>
      </Toolbar>
    </AppBar>
  );
};
