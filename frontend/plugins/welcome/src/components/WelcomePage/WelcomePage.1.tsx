import React from 'react';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import CloseIcon from '@material-ui/icons/Close';
import Alert from '@material-ui/lab/Alert';
import Collapse from '@material-ui/core/Collapse';
import MenuIcon from '@material-ui/icons/Menu';


import {
  createMuiTheme,
  createStyles,
  makeStyles,
  Theme,
  ThemeProvider,
} from '@material-ui/core/styles';

const theme = createMuiTheme({
  palette: {
    primary:{
      main: '#f44336',
    } ,
    secondary: {
      main: '#81c784',
    },
  },
});




const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    
    textField: {
      marginLeft: theme.spacing(1),
      marginRight: theme.spacing(1),
      width: 200,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
    },
    selectEmpty: {
      marginTop: theme.spacing(2),
    },
    paper2: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: 'white',
      background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
    },
    paper3: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: 'white',
      background: 'linear-gradient(45deg, #ffe082 30%, #ff8f00 90%)',
    },
  }),
);

export default function WelcomePage() {
  const classes = useStyles();
  const [age, setAge] = React.useState('');

  const handleChange = (event: React.ChangeEvent<{ value: unknown; }>) => {
    setAge(event.target.value as string);
  };
  const [open, setOpen] = React.useState(false);

  return (
    <div>

      <AppBar position="static">
        <Toolbar>
          <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" className={classes.title}>
            News
          </Typography>
          <FormControl className={classes.formControl}>
        <InputLabel id="demo-simple-select-label">บัญชีผู้ใช้</InputLabel>
        <Box color="text.primary" clone>
        <Button />
      </Box>
        <Select
          labelId="demo-simple-select-label"
          id="demo-simple-select"
          value={age}
          onChange={handleChange}
        >
          <MenuItem value={10}>patza30596@gmail.com</MenuItem>
          <MenuItem value={20}>ไทยพานิช</MenuItem>
          <MenuItem value={30}>กสิกรไทย</MenuItem>
          <MenuItem value={30}>กรุงศรี</MenuItem>
          <MenuItem value={30}>ออมสิน</MenuItem>
        </Select>
        </FormControl>
         </Toolbar>
      </AppBar>

    
      <Paper variant="outlined" elevation={3} className={classes.paper3}>
      <div>

      <Grid container spacing={4}>
        <Grid item xs={12}>
          </Grid>
        <Grid item xs={1}>
          </Grid>
        <Grid item xs={2} >
          <Paper elevation={3} className={classes.paper2}>
            <Typography variant="h4" gutterBottom>
              การชำระเงิน
            </Typography>
          </Paper>
        </Grid>
        <Grid item xs={9}>
          </Grid>
          <Grid item xs={1}>
          </Grid>
        <Grid item xs={2} >
        <TextField
          id="outlined-read-only-input"
          label="หมายเลขใบการจอง"
          defaultValue="001"
          InputProps={{
            readOnly: true,
          }}
          variant="outlined"
      ></TextField>
        </Grid>
        <Grid item xs={9} ></Grid>
          <Grid item xs={1}>
          </Grid>
        <Grid item xs={2} >
        <TextField
          id="outlined-read-only-input"
          label="จำนวนเงินที่ต้องชำระ"
          defaultValue="300"
          InputProps={{
            readOnly: true,
          }}
          helperText="บาท"
          variant="outlined"
      ></TextField>
        </Grid>
        <Grid item xs={5} >
          
        </Grid>
      </Grid>

      </div>
      

      

      <div>
      <Grid container spacing={4}>
      <Grid item xs={1} >
        </Grid>
        <Grid item xs={2} >
        <FormControl className={classes.formControl}>
        <InputLabel id="demo-simple-select-label">ธนาคาร</InputLabel>
        <Box color="text.primary" clone>
        <Button />
      </Box>
        <Select
          labelId="demo-simple-select-label"
          id="demo-simple-select"
          value={age}
          onChange={handleChange}
        >
          <MenuItem value={10}>กรุงไทย</MenuItem>
          <MenuItem value={20}>ไทยพานิช</MenuItem>
          <MenuItem value={30}>กสิกรไทย</MenuItem>
          <MenuItem value={30}>กรุงศรี</MenuItem>
          <MenuItem value={30}>ออมสิน</MenuItem>
        </Select>
        </FormControl>
        </Grid>
        <Grid item xs={9} >
        </Grid>
        </Grid>
        </div>

      <div>
      <Grid container spacing={4}>
      <Grid item xs={1} >
        </Grid>
        <Grid item xs={2} >
        <TextField id="outlined-basic" label="เลขอ้างอิง"  variant="outlined" />
        </Grid>
        <Grid item xs={9} >
        </Grid>
        </Grid>
      </div>
      


      <div>
      <Grid container spacing={4}>
      <Grid item xs={1} >
        </Grid>
        <Grid item xs={2} >
        <TextField
        id="datetime-local"
        label="วันเวลาการโอน"
        type="datetime-local"
        defaultValue="2017-05-24T10:30"
        className={classes.textField}
        InputLabelProps={{
          shrink: true,
        }}
      />
        </Grid>
        <Grid item xs={9} >
        </Grid>
        </Grid>
      </div>


      <div>
      <Grid container spacing={4}>
      <Grid item xs={1} >
        </Grid>
        <Grid item xs={1} >
        <ThemeProvider theme={theme}>
      <Button
        disabled={false}
        variant="contained"
        color="secondary"
        onClick={() => {
          setOpen(true);
        }}
      >
        ยืนยัน
      </Button>
      </ThemeProvider>
        </Grid>
        <Grid item xs={1} >
        <ThemeProvider theme={theme}>
        <Button variant="contained" color="primary">
        ยกเลิก
      </Button>
      </ThemeProvider>
        </Grid>
        <Grid item xs={1} ></Grid>
        <Grid item xs={3} >
          <Collapse in ={open}>
        <Alert
          action={
            <IconButton
              aria-label="close"
              color="inherit"
              size="small"
              onClick={() => {
                setOpen(false);
              }}
            >
              <CloseIcon fontSize="inherit" />
            </IconButton>
          }
        >
          บันทึกสำเร็จเเล้ว
        </Alert>
      </Collapse></Grid>
        </Grid>
      </div>
      
      </Paper>
    </div>
  );
}
