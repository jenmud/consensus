//import { makeStyles } from "@mui/material";
import styled from "@emotion/styled";
import { AppBar, Box, InputBase, List, ListItem, ListItemButton, ListItemIcon, ListItemText, Stack, Switch, Toolbar, Typography } from "@mui/material";
import HandshakeIcon from '@mui/icons-material/Handshake';
import HomeIcon from '@mui/icons-material/Home';
import AccountTreeIcon from '@mui/icons-material/AccountTree';
import PlumbingIcon from '@mui/icons-material/Plumbing';
import FeaturedPlayListIcon from '@mui/icons-material/FeaturedPlayList';
import Brightness4Icon from '@mui/icons-material/Brightness4';
import React from "react";
import DraggableCard from "./draggable";

//const useStyles = makeStyles((theme) => ({
//  root: {
//    flexGrow: 1,
//  },
//  menuButton: {
//    marginRight: theme.spacing(2),
//  },
//}));

const SideBar = () => {
	return (
		<Box flex={1} p={2} sx={{ display: { xs: "none", sm: "block"}}}>
      <List>
        <ListItem disablePadding>
          <ListItemButton component="a" href="/projects">
            <ListItemIcon><HomeIcon/></ListItemIcon>
            <ListItemText primary="Home"/>
          </ListItemButton>
        </ListItem>
        <ListItem disablePadding>
          <ListItemButton component="a" href="/projects">
            <ListItemIcon><AccountTreeIcon/></ListItemIcon>
            <ListItemText primary="Projects"/>
          </ListItemButton>
        </ListItem>
        <ListItem disablePadding>
          <ListItemButton component="a" href="/epics">
            <ListItemIcon><PlumbingIcon/></ListItemIcon>
            <ListItemText primary="Epics"/>
          </ListItemButton>
        </ListItem>
        <ListItem disablePadding>
          <ListItemButton component="a" href="/features">
            <ListItemIcon><FeaturedPlayListIcon/></ListItemIcon>
            <ListItemText primary="Features"/>
          </ListItemButton>
        </ListItem>
        <ListItem disablePadding>
          <ListItemButton component="a" href="/features" disableRipple>
            <ListItemIcon><Brightness4Icon/></ListItemIcon>
            <Switch></Switch>
          </ListItemButton>
        </ListItem>
      </List>
    </Box>
	)
}

const RightBar = () => {
  return (
    <Box bgcolor={"lightblue"} flex={4} p={2} sx={{ display: { xs: "none", sm: "block"}}}>
      <DraggableCard/>
    </Box>
  )
}

const Feed = () => {
  return (
    <Box bgcolor={"lightsalmon"} flex={2} p={2}>feed</Box>
  )
}

const StyledToolbar = styled(Toolbar)({
  display:"flex",
  justifyContent:"space-between"
})

const Search = styled(Box)(({theme})=>({
  backgroundColor:"white",
  padding:"0 20px",
  borderRadius: theme.shape.borderRadius,
  width: "20%"
}))

export const ConcensusNavBar = () => {

  return (
    <Box>
    <AppBar position="sticky">
      <StyledToolbar>
        <Typography variant="h6" sx={{display:{xs:"none", sm: "block"}}}>Concensus</Typography>
        <HandshakeIcon sx={{display:{xs:"block", sm: "none"}}}/>
        <Search><InputBase placeholder="search"></InputBase></Search>
      </StyledToolbar>
    </AppBar>
    <Stack direction={"row"} spacing={2} justifyContent={"space-between"}>
      <SideBar/>
      <RightBar/>
      <Feed/>
    </Stack>
    </Box>
  );
};
