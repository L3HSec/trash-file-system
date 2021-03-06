import React from 'react';
import { Dialog, DialogTitle, Button, makeStyles } from '@material-ui/core'
import { Grid, TextField } from '@material-ui/core';

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
  };
  return (
    <Dialog open={open} onClose={handleClose} >
      <DialogTitle id="simple-dialog-title" className={classes.dialog}>{fileInfo.FileName}</DialogTitle>
      <Grid container direction="column">
        <Grid item xs className={classes.grid}>
          <Grid item xs={12}>
            <TextField multiline variant="outlined" rows="10" className={classes.comment} value={fileInfo.Comment}></TextField>
          </Grid>
          <Grid item xs={12}>
            <a href={"/file/"+fileInfo.FileID}>
              <Button variant="outlined" className={classes.download} >下载</Button>
            </a>
          </Grid>

        </Grid>
      </Grid>

    </Dialog>

  );
}
