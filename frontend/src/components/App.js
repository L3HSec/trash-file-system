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
          <HeaderBar />

          <FileTable loadedCallback={this.refreshFileList} />
        </ThemeProvider>
    );
  }

}

export default App;
