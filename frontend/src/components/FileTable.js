import React, { Component } from 'react';
import PropTypes from 'prop-types';
import {Button, Container, Paper, Box} from '@material-ui/core';
import {Table, TableHead, TableRow, TableCell,TableBody,  TableContainer } from '@material-ui/core';
import {ArrowDownwardOutlined} from '@material-ui/icons'
import axios from 'axios';

class FileItem extends Component {
  static propTypes = {
    fileInfo: PropTypes.object,
    clickHandler: PropTypes.func
  };

  render() {
    return (
      // <tr key={this.props.fileInfo.itemId} className="fileItem">
      //   <td>{this.props.fileInfo.fileName}</td>
      //   <td>{this.props.fileInfo.fileSize}</td>
      //   <td>{this.props.fileInfo.fileRemainingTime}</td>
      //   <td><Button >下载</Button></td>
      // </tr>
      <TableRow key={this.props.fileInfo.itemId}>
        <TableCell>{this.props.fileInfo.fileName}</TableCell>
        <TableCell>{this.props.fileInfo.fileSize}</TableCell>
        <TableCell>{this.props.fileInfo.fileRemainingTime}</TableCell>
      </TableRow>

    );
  };
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
            fileSize: "1KB",
            fileRemainingTime: "+10000s"
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
      return <div>Loading</div>
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
                  <TableCell>剩余过期时间</TableCell>
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