import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Avatar from '@material-ui/core/Avatar';
import CarIcon from '@material-ui/icons/DirectionsCar';
import Logo from './Logo';

const styles = theme => ({
  root: {
    position: 'relative',
    overflow: 'auto',
    maxHeight: '85%',
    backgroundColor: theme.palette.background.paper,
  },
});

function FolderList(props) {
  const { classes } = props;
  return (
    <div className={classes.root}>
    { props.data.length > 0 ?
      <List>
          {props.data.map((passport) => {
              return(
            <ListItem key={passport.Key}>
                <Avatar>
                  <CarIcon />
                </Avatar>
                {passport.Record ? <ListItemText primary={passport.Key} secondary={`${passport.Record.fname} ${passport.Record.gender} ${passport.Record.dob} ${passport.Record.place} (${passport.Record.status})`} />
                : <ListItemText primary={passport.Key} secondary={passport.Msg}/>
                }
            </ListItem>
          )})}
      </List>
    : <Logo />
    }
    </div>
  );
}

FolderList.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(FolderList);