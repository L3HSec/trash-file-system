import React from 'react';
import { Dialog, TextField, Grid, DialogTitle, Button, makeStyles } from '@material-ui/core';
import {useSnackbar} from 'notistack';
import axios from 'axios';

const useStyles = makeStyles(theme => ({
  dialog: {
    minWidth: "30vw",
  },
  grid: {
    textAlign: "center",
  },
  comment: {
    minWidth: "30vw",
    minHeight: "35vh",
  },
  input: {
    display: "none",
  },
  selectFile: {
    minWidth: "28%",
    margin: "4vh",
  },
  upload: {
    minWidth: "28%",
    margin: "4vh",
  }
}));

export default function UploadDialog(props) {
  const { open, onClose } = props;
  const { enqueueSnackbar } = useSnackbar();
  const classes = useStyles();


  var fileInfo = {}
  const handleClose = () => {
    onClose();
  }

  const selectFile = (e) => {
    fileInfo.file = e.target.files[0];
  }

  const updateComment = (e) => {
    fileInfo.comment=e.target.value;
  }

  const doUpload = () => {
    console.log(fileInfo);
    if (!fileInfo.file) {
      alert("请选择文件");
      return;
    }
    var param = new FormData();
    param.append('file',fileInfo.file);
    param.append('comment', fileInfo.comment);
    var config = {
      headers: {'Content-Type': 'multipart/form-data'}
    }
    axios.post("/file",param, config)
    .then(res => {
      //todo
      console.log(res);
      if (res.status === 201) {
        enqueueSnackbar("上传成功",{variant: "success"});
        onClose();
      }
    })
    .catch(error => {
      enqueueSnackbar("上传失败",{variant: "success"});
    });
  }

  return (
    <Dialog open={open} onClose={handleClose} >
      <DialogTitle className={classes.dialog}>上传文件</DialogTitle>
      <Grid container direction="column">
        <Grid item xs={12} className={classes.grid}>
          <TextField multiline variant="outlined" rows="10" className={classes.comment} onChange={updateComment}></TextField>
        </Grid>
        <Grid item xs={12} className={classes.grid}>
          <input
            accept="*"
            className={classes.input}
            type="file"
            id="text-button-file"
            onChange={selectFile}
          />
          <label htmlFor="text-button-file">
            <Button component="span" variant="contained" color="primary" className={classes.selectFile}>选择文件</Button>
          </label>
          <Button variant="contained" color="secondary" className={classes.upload} onClick={doUpload}>上传</Button>
        </Grid>
      </Grid>
    </Dialog>
  
  );
}