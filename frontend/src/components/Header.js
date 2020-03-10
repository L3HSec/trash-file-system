import React from 'react';
import { AppBar, Toolbar, Typography, Button } from '@material-ui/core';
import { makeStyles } from '@material-ui/core';

const useStyles = makeStyles(theme => ({
  title: {
    flexGrow: 1
  }
}))

export default function HeaderBar() {
  const classes = useStyles();

  return (
    <AppBar position="static">
      <Toolbar>
        <Typography variant="h6" className={classes.title}>
          L3H Cloud
        </Typography>
        <Button color="inherit">
          上传文件
        </Button>
      </Toolbar>
    </AppBar>
  );
}