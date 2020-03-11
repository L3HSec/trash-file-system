import React, { Component } from 'react';
import axios from 'axios';
import PropTypes from 'prop-types';
import { dateFormat, getfilesize } from './DataFormat';
import { withSnackbar } from 'notistack';

import { Container, Paper, Box } from '@material-ui/core';
import { Table, TableHead, TableRow, TableCell, TableBody, TableContainer } from '@material-ui/core';

import DownloadDialog from './FileDownloadDialog';

function FileItem(props) {
  const [open, setOpen] = React.useState(false);

  const handleClickFile = () => {
    setOpen(true);
  }

  const handleClose = () => {
    setOpen(false);
  }

  var date = dateFormat("YYYY-mm-dd HH:MM", new Date(props.fileInfo.Expire*1000));
  var fileSize = getfilesize(props.fileInfo.FileSize);

  return (
    <React.Fragment>
      <DownloadDialog open={open} onClose={handleClose} fileInfo={props.fileInfo} />
      <TableRow key={props.fileInfo.FileID} onClick={() => { handleClickFile() }}>
        <TableCell style={{ width: "60%" }}>{props.fileInfo.FileName}</TableCell>
        <TableCell>{fileSize}</TableCell>
        <TableCell>{date}</TableCell>
      </TableRow>

    </React.Fragment>
  );
}

class FileTable extends Component {
  // filelist = [];
  constructor(props) {
    super(props);
    this.state = {
      filelist: [],
      isLoaded: false
    };
  };

  static propTypes = {
    loadedCallback: PropTypes.func,
    enqueueSnackbar: PropTypes.func
  }

  componentDidMount() {
    this.refreshFileList();
  }
  refreshFileList(props) {
    const _this = this;
    axios.get('/file')
      .then((response) => {
        // response = [
        //   {
        //     itemId: "asd",
        //     fileName: "test",
        //     fileSize: 111111,
        //     fileRemainingTime: 1583859517598
        //   }
        // ]
        // response = JSON.stringify(response);
        _this.setState({
          filelist: response.data,
          isLoaded: true
        });
        _this.renderTableData();
        this.props.enqueueSnackbar("获取文件列表成功",{variant: "success"});
      })
      .catch((error) => {
        console.log(error);
        this.props.enqueueSnackbar("获取文件列表失败",{variant: "error"});
      });
  };

  renderTableData() {
    return this.state.filelist.map((file, index) => {
      return (
        <FileItem fileInfo={file} />
      )

    });
  }

  render() {
    if (!this.state.isLoaded) {
      return <center><div>Loading</div></center>
    } else {
      return (
        <Container>
          <Box component="span" display="block" height="30px" > </Box>

          <TableContainer component={Paper}>
            <Table className="filetable">
              <TableHead>
                <TableRow>
                  <TableCell style={{ width: "60%" }}>文件名</TableCell>
                  <TableCell>大小</TableCell>
                  <TableCell>到期时间</TableCell>
                </TableRow>
              </TableHead>
              <TableBody stripedRows>
                {this.renderTableData()}
              </TableBody>
            </Table>
          </TableContainer>
        </Container>

      );
    }


  }

};

export default withSnackbar(FileTable);