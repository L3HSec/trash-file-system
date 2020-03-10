import React from 'react';
import { Dialog, DialogTitle, Button, Box, makeStyles } from '@material-ui/core'
import { Grid, TextField, ClickAwayListener } from '@material-ui/core';

const useStyles = makeStyles(theme => ({
  dialog: {
    minWidth: "30vw",
  },
  grid: {
    textAlign: "center"
  },
  comment: {
    minWidth: "30vw",
    minHeight: "35vh",
  },
  download: {
    minWidth: "28%",
    margin: "4vh"
  }
}));
export default function DownloadDialog(props) {
  const { open, fileInfo, onClose } = props;
  const classes = useStyles();
  const handleClose = () => {
    onClose();
  }
  return (
    // <ClickAwayListener onClickAway={handleClose}>
    <Dialog open={open} onClose={handleClose} >
      <DialogTitle id="simple-dialog-title" className={classes.dialog}>{fileInfo.fileName}</DialogTitle>
      <Grid container direction="column">
        <Grid item xs className={classes.grid}>
          <Grid item xs={12}>
            <TextField multiline variant="outlined" rows="10" className={classes.comment} value={fileInfo.comment}></TextField>
          </Grid>
          <Grid item xs={12}>
            <Button variant="outlined" className={classes.download} >下载</Button>
          </Grid>

        </Grid>
      </Grid>

    </Dialog>

    // </ClickAwayListener>
  );
}
