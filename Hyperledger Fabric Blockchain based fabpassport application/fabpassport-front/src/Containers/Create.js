import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

const styles = theme => ({
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    width: 200,
  },
  menu: {
    width: 200,
  },
});

class Create extends React.Component {
  state = {
    ID: null,
    dob: null,
    fname: null,
    gender: null,
    status: null
  };

  handleChange = name => event => {
    this.setState({
      [name]: event.target.value,
    });
  };

  createHandler = () => {
    //Check form validity
    if (!(this.state.ID && this.state.fname && this.state.gender && this.state.dob && this.state.status)){
      alert('All fields must be filled in');
    } else if (this.state.ID.slice(0,8) !== 'PASSPORT') {
        alert('ID MUST CONTAIN "PASSPORT" FOLLOWED BY ID')
    } else if (this.state.ID.slice(8).length > 8 || isNaN(this.state.ID.slice(8))) {
        alert('ID MUST CONTAIN "PASSPORT" FOLLOWED BY ID BETWEEN 0 AND 999')
    } else {
      this.props.switchFeedHandler(1)
      this.props.socket.emit('REQUEST', {action: "CREATE", data:this.state})
    }
  }

  render() {
    const { classes } = this.props;

    return (
      <form className="Main-inside" noValidate autoComplete="off">
        <Typography  variant="display2">
          Create a passport
        </Typography>
        <TextField
          label="PASSPORT ID"
          className={classes.textField}
          value={this.state.name}
          onChange={this.handleChange('ID')}
          margin="normal"
        />
        <TextField
          label="Dob"
          className={classes.textField}
          value={this.state.name}
          onChange={this.handleChange('dob')}
          margin="normal"
        />
        <TextField
          label="Fname"
          className={classes.textField}
          value={this.state.name}
          onChange={this.handleChange('fname')}
          margin="normal"
        />
        <TextField
          label="Gender"
          className={classes.textField}
          value={this.state.name}
          onChange={this.handleChange('gender')}
          margin="normal"
        />
        <TextField
          label="Status"
          className={classes.textField}
          value={this.state.name}
          onChange={this.handleChange('status')}
          margin="normal"
        />
        <Button variant="contained" 
                color="primary" 
                disabled={!this.props.connected}
                className={classes.button} 
                onClick={this.createHandler}>
           {this.props.connected ? "CREATE" : "DISCONNECTED"}
        </Button>
        <p>PASSPORT ID is case sensitive and should start with 'PASSPORT' followed by digits (e.g. PASSPORT10)</p>
      </form>
      
    );
  }
}


export default withStyles(styles)(Create);