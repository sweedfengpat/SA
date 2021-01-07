import React from 'react';
import {useEffect,useState } from 'react';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';

import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import InputAdornment from '@material-ui/core/InputAdornment';
import Swal from 'sweetalert2'; // alert

import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntUser } from '../../api/models/EntUser'; // import interface User
import { EntConfirmation } from '../../api/models/EntConfirmation'; // import interface Confirmation
import { EntBank } from '../../api/models/EntBank'; // import interface Bank





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


  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
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
    container: {
      display: 'flex',
      flexWrap: 'wrap',
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
    formControl: {
      margin: theme.spacing(0),
      minWidth: 220,
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

interface savebill {
  paynumber:    number;
	paytotal:     number;
	bank:         number;
	confirmation: number;
  user:    number;
  paytime:      string;
}

export default function WelcomePage() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [bill, setBill] = React.useState<Partial<savebill>>({});
  const [users, setUsers] = React.useState<EntUser[]>([]);
  const [confirmations, setConfirmations] = React.useState<EntConfirmation[]>([]);
  const [banks, setBanks] = React.useState<EntBank[]>([]);

  const getBank = async () => {
    const res = await http.listBank({ limit: 10, offset: 0 });
    setBanks(res);
  };
  const getUsers = async () => {
    const res = await http.listUser({ limit: 10, offset: 0 });
    setUsers(res);
  };

  const getConfirmation = async () => {
    const res = await http.listConfirmation({ limit: 10, offset: 0 });
    setConfirmations(res);
  };
 
  useEffect(() => {
    getUsers();
    getConfirmation();
    getBank();
  }, []);

  const handleChange = (event: React.ChangeEvent<{ name?: string; value: unknown; }>,
    ) => {
      const name = event.target.name as keyof typeof WelcomePage;
      const { value } = event.target;
      setBill({ ...bill, [name]: value });
      console.log(bill);
  };
  const handleChange2 = (event: React.ChangeEvent<{ name?: string; value: unknown; }>,
    ) => {
      const name = event.target.name as keyof typeof WelcomePage;
      const { value } = event.target;
      setBill({ ...bill, [name]: Number(value) });
      console.log(bill);
  };
  const handleChange3 = (event: React.ChangeEvent<{ name?: string; value: unknown; }>,
    ) => {
      const name = event.target.name as keyof typeof WelcomePage;
      const { value } = event.target;
      setBill({ ...bill, [name]: Number(value)/10 | 0 });
      console.log(bill);
  };

  function clear() {
    setBill({});
  };

  
  // function save data
  function save() {
    bill.paytime  +=":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/bills';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(bill),
    };

    console.log(bill); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console
    fetch(apiUrl, requestOptions)
      .then(response => {
        console.log(response)
        response.json()
        if (response.ok === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      })
  }
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
          <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>สมาชิก</InputLabel>
                <Select
                  value={bill.user || ''}
                  onChange={handleChange}
                  name="user"
                >
                  {users.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id} >
                        {item.email}
                      </MenuItem>
                    );
                  })}
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
          <Grid item xs={1} >
        </Grid>
        <Grid item xs={2} >
        <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>ใบจอง</InputLabel>
                <Select
                  value={bill.confirmation || ''}
                  onChange={handleChange}
                  name="confirmation"
                >
                  {confirmations.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id} >
                        {item.id}
                      </MenuItem>
                    );
                  })}

                  
                </Select>
              </FormControl>
        </Grid>
        <Grid item xs={9} >
        </Grid>
        <Grid item xs={1} >
        </Grid>
        <Grid item xs={2} >
        <FormControl variant="outlined" className={classes.formControl}>
        <TextField
          id="paytotal"
          name="paytotal"
          label="จำนวนเงินที่ต้องชำระ"
          defaultValue = {0}
          value={Number(confirmations.filter(pay => pay.id === bill.confirmation).map(item => {
               return Number((item.hourstime))
          }))*50}
          onChange={handleChange3}
          InputProps={{
            //readOnly: true,
            endAdornment: <InputAdornment position="end">Bath</InputAdornment>,
          }}
          variant="outlined"
      ></TextField>
        </FormControl>
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
        <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>ธนาคาร</InputLabel>
                <Select
                  value={bill.bank || ''}
                  onChange={handleChange}
                  name="bank"
                >
                  {banks.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.bank}
                      </MenuItem>
                    );
                  })}
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
        <FormControl variant="outlined" className={classes.formControl}>
        <TextField id="outlined-basic" name="paynumber" label="เลขอ้างอิง"  variant="outlined" value={bill.paynumber} onChange={handleChange2}/>
        </FormControl>
        </Grid>
        <Grid item xs={9} >
        </Grid>
        </Grid>
      </div>
      


      <div>
      <Grid container spacing={4} justify="space-around" >
      <Grid item xs={1} >
        </Grid>
        <Grid item xs={2} >
        <FormControl variant="outlined" className={classes.formControl}>
        <form className={classes.container} noValidate>
                <TextField
                  label="เวลาโอน"
                  name="paytime"
                  type="datetime-local"
                  value={bill.paytime}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
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
        <Grid item xs={1} >
        <ThemeProvider theme={theme}>
      <Button
        disabled={false}
        variant="contained"
        color="secondary"
        onClick={() => {
          save();
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
        
        </Grid>
      </div>
      
      </Paper>
    </div>
  );
}
