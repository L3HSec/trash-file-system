import React from 'react';
import { AppBar, Toolbar, Typography, Button } from '@material-ui/core';
import { makeStyles } from '@material-ui/core';

import UploadDialog from './FileUploadDialog';

const useStyles = makeStyles(theme => ({
  title: {
    flexGrow: 1
  }
}));

export default function HeaderBar(props) {
  const classes = useStyles();
  const [open, setOpen] = React.useState(false);

  const showUploadDialog = () => {
    setOpen(true);
  }
  const closeDialog = () => {
    setOpen(false);
    props.refreshFileList();
  }

  return (
    <AppBar position="static">
      <Toolbar>
        <Typography variant="h6" className={classes.title}>
          L3H Cloud
        </Typography>
        <Button color="inherit" onClick={showUploadDialog}>
          Upload
        </Button>
        <UploadDialog open={open} onClose={closeDialog} />
      </Toolbar>
    </AppBar>
  );
}