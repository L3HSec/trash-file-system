import React, { Component } from 'react';
import FileTable from './FileTable';
import HeaderBar from './Header';
import {createMuiTheme, ThemeProvider} from '@material-ui/core/styles';
import blue from '@material-ui/core/colors/blue';

const theme = createMuiTheme({
  palette: {
    primary: blue
  }
})

class App extends Component {

  render() {
    return (
        <ThemeProvider theme={theme}>
          <HeaderBar refreshFileList={() => this.content.refreshFileList()} />

          <FileTable loadedCallback={this.refreshFileList} ref={instance=>{this.content=instance}} />
        </ThemeProvider>
    );
  }

}

export default App;
