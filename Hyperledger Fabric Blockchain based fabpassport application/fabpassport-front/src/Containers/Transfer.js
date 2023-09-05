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

class Transfer extends React.Component {
  state = {

  };

  handleChange = name => event => {
    this.setState({
      [name]: event.target.value,
    });
  };

  submitHandler = () => {
    //check if input is correctly formatted
    if (!(this.state.ID && this.state.newStatus)){
      alert('All fields must be filled in');
    } else {
      this.props.switchFeedHandler(1)
      this.props.socket.emit('REQUEST', {action: "CHANGE", data:this.state})
    }
  }

  render() {
    const { classes } = this.props;

    return (
      <form className="Main-inside" noValidate autoComplete="off">
     <Typography  variant="display2">
      Change a status
      </Typography>
      <TextField
          label="PASSPORT ID"
          className={classes.textField}
          value={this.state.name}
          onChange={this.handleChange('ID')}
          margin="normal"
        />
        <TextField
          label="New Status"
          className={classes.textField}
          value={this.state.name}
          onChange={this.handleChange('newStatus')}
          margin="normal"
        />
        <Button variant="contained" color="primary" disabled={!this.props.connected} className={classes.button} onClick={this.submitHandler}>
            {this.props.connected ? "CHANGE" : "DISCONNECTED"}
        </Button>
        <p>Passport ID is case sensitive must be valid</p>
      </form>
      
    );
  }
}


export default withStyles(styles)(Transfer);