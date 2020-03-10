import React, { Component } from 'react';
import axios from 'axios';
import PropTypes from 'prop-types';
import {dateFormat, getfilesize} from './DataFormat';

import {Container, Paper, Box } from '@material-ui/core';
import {Table, TableHead, TableRow, TableCell,TableBody,  TableContainer } from '@material-ui/core';

import DownloadDialog from './FileDownloadDialog';

function FileItem (props){
  const [open, setOpen] = React.useState(false);

  const handleClickFile = () => {
    setOpen(true);
  }

  const handleClose = () => {
    setOpen(false);
    console.log("aaaa");
  }

  var date = dateFormat("YYYY-mm-dd HH:MM",new Date(props.fileInfo.fileRemainingTime));
  var fileSize = getfilesize(props.fileInfo.fileSize);
  
  return (
    <TableRow key={props.fileInfo.itemId} onClick={()=>{handleClickFile()}}>
      <TableCell>{props.fileInfo.fileName}</TableCell>
      <TableCell>{fileSize}</TableCell>
      <TableCell>{date}</TableCell>
      <DownloadDialog open={open} onClose={handleClose} fileInfo={props.fileInfo} />
    </TableRow>
  );
}

export default class FileTable extends Component {
  // filelist = [];
  constructor(props) {
    super(props);
    this.state = {
      filelist: [],
      isLoaded: false
    };
  };

  static propTypes = {
    loadedCallback: PropTypes.func
  }

  componentDidMount() {
    this.refreshFileList();
  }
  refreshFileList(props) {
    const _this = this;
    axios.get('/file')
      .then((response) => {
        response = [
          {
            itemId: "asd",
            fileName: "test",
            fileSize: 111111,
            fileRemainingTime: 1583859517598
          }
        ]
        response = JSON.stringify(response);
        _this.setState({
          filelist: JSON.parse(response),
          isLoaded: true
        });
        _this.renderTableData();
      })
      .catch((error) => {
        console.log(error);
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
                  <TableCell style={{width: "60%"}}>文件名</TableCell>
                  <TableCell>大小</TableCell>
                  <TableCell>到期时间</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {this.renderTableData()}
              </TableBody>
            </Table>
          </TableContainer>
        </Container>

      );
    }


  }

}