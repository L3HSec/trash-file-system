import React from 'react';
import { Dialog, TextField, Box, Grid, DialogTitle, Button, makeStyles } from '@material-ui/core';
import {useState} from 'react';
import {useSnackbar} from 'notistack';
import axios from 'axios';

const useStyles = makeStyles(theme => ({
  dialog_banner: {
    minWidth: "30vw",
  },
  dialog: {
    minWidth: "30vw",
  },
  grid: {
    textAlign: "center",
  },
  buttons: {
    textAlign: "center",
    maxHeight: "13vh"
  },
  comment: {
    minWidth: "30vw",
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
  },
  filename: {
    minWidth: "30vw",
    margin: "1vh",
    textAlign: "center"
  }
}));

export default function UploadDialog(props) {
  const { open, onClose } = props;
  const { enqueueSnackbar } = useSnackbar();
  const classes = useStyles();
  const [name, setName] = useState('');
  const [fileinfo, setFileinfo] = useState({});
  const [filecomment, setFilecomment] = useState('');

  const handleClose = () => {
    onClose();
  }

  const selectFile = (e) => {
    setFileinfo(e.target.files[0])
    setName("选择的文件: "+e.target.files[0].name);
  }

  const updateComment = (e) => {
    setFilecomment(e.target.value);
    // fileInfo.comment=e.target.value;
  }

  const doUpload = () => {
    console.log(fileinfo);
    if (fileinfo == {}) {
      alert("请选择文件");
      return;
    }
    var param = new FormData();
    param.append('file',fileinfo);
    param.append('comment', filecomment);
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
    setName('');
  }

  return (
    <Dialog open={open} onClose={handleClose} className={classes.dialog} >
      <DialogTitle className={classes.dialog_banner}>上传文件</DialogTitle>
      <Grid container direction="column">
        <Grid item xs={12} className={classes.grid}>
          <TextField multiline variant="outlined" rows="10" className={classes.comment} onChange={updateComment}></TextField>
        </Grid>
        <Grid item xs={12} className={classes.buttons}>
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
        <Grid item xs={12} className={classes.filename}>
          <Box className={classes.filename}>{name}</Box>
        </Grid>
      </Grid>
    </Dialog>
  
  );
}