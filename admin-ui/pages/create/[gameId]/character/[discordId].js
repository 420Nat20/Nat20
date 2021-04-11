import React from "react";
import { useRouter } from "next/router";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import CssBaseline from "@material-ui/core/CssBaseline";
import TextField from "@material-ui/core/TextField";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Checkbox from "@material-ui/core/Checkbox";
import Link from "@material-ui/core/Link";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import { makeStyles } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";
import { useForm, Controller } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import AlertDialog from "../../../../components/dialog";

const schema = yup.object().shape({
  PlayerName: yup.string().required(),
  Name: yup.string().required(),
  Class: yup.string().required(),
  Background: yup.string().required(),
  Race: yup.string().required(),
  Alignment: yup.string().required(),
});

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {"Copyright © "}
      <Link color="inherit" href="https://material-ui.com/">
        Nat20 - HackKU
      </Link>{" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
}

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: "100%", // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

export default function CharacterCreate() {
  const classes = useStyles();
  const {
    control,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();
  const [open, setOpen] = React.useState(false);

  const router = useRouter();
  const { gameId, discordId } = router.query;

  const onSubmit = async (data) => {
    const reqBody = {
      DiscordID: +discordId,
      Name: data.Name,
      PlayerName: data.PlayerName,
      Class: data.Class,
      Background: data.Background,
      Race: data.Race,
      Alignment: data.Alignment,
      Strength: +data.Strength,
      Dexterity: +data.Dexterity,
      Constitution: +data.Constitution,
      Intelligence: +data.Intelligence,
      Wisdom: +data.Wisdom,
      Charisma: +data.Charisma,
    };

    const response = await fetch(
      `https://api.nat20.tech/games/${gameId}/users/`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(reqBody),
      }
    );

    if ((await response.status) === 200) {
      reset();
      setOpen(true);
    }
  };

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      {!open ? (
        <div className={classes.paper}>
          <Typography component="h1" variant="h5">
            Character Creation
          </Typography>
          <form className={classes.form} onSubmit={handleSubmit(onSubmit)}>
            <Controller
              as={TextField}
              name="PlayerName"
              control={control}
              render={({ field }) => (
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  label="Player Name"
                  id="player"
                  {...field}
                />
              )}
            />

            {/* Character Info */}
            <Controller
              name="Name"
              control={control}
              render={({ field }) => (
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  id="name"
                  label="Character Name"
                  name="Name"
                  {...field}
                />
              )}
            />
            <Controller
              name="Class"
              control={control}
              render={({ field }) => (
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  name="Class"
                  label="Class"
                  id="class"
                  defaultValue=""
                  {...field}
                />
              )}
            />
            <Controller
              name="Background"
              control={control}
              render={({ field }) => (
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  name="Background"
                  label="Background"
                  id="background"
                  {...field}
                />
              )}
            />
            <Controller
              name="Race"
              control={control}
              render={({ field }) => (
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  name="Race"
                  label="Race"
                  id="race"
                  {...field}
                />
              )}
            />
            <Controller
              name="Alignment"
              control={control}
              render={({ field }) => (
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  name="Alignment"
                  label="Alignment"
                  id="alignment"
                  {...field}
                />
              )}
            />

            {/* Main Stats */}
            <Grid container spacing={2}>
              <Grid item xs={6}>
                <Controller
                  name="Strength"
                  control={control}
                  render={({ field }) => (
                    <TextField
                      variant="outlined"
                      margin="normal"
                      required
                      fullWidth
                      name="Strength"
                      label="Strength"
                      id="strength"
                      type="number"
                      {...field}
                    />
                  )}
                />
              </Grid>
              <Grid item xs={6}>
                <Controller
                  name="Dexterity"
                  control={control}
                  render={({ field }) => (
                    <TextField
                      variant="outlined"
                      margin="normal"
                      required
                      fullWidth
                      name="Dexterity"
                      label="Dexterity"
                      id="dexterity"
                      type="number"
                      {...field}
                    />
                  )}
                />
              </Grid>
              <Grid item xs={6}>
                <Controller
                  name="Constitution"
                  control={control}
                  render={({ field }) => (
                    <TextField
                      variant="outlined"
                      margin="normal"
                      required
                      fullWidth
                      name="Constitution"
                      label="Constitution"
                      id="constitution"
                      type="number"
                      {...field}
                    />
                  )}
                />
              </Grid>
              <Grid item xs={6}>
                <Controller
                  name="Intelligence"
                  control={control}
                  render={({ field }) => (
                    <TextField
                      variant="outlined"
                      margin="normal"
                      required
                      fullWidth
                      name="Intelligence"
                      label="Intelligence"
                      id="intelligence"
                      type="number"
                      {...field}
                    />
                  )}
                />
              </Grid>
              <Grid item xs={6}>
                <Controller
                  name="Wisdom"
                  control={control}
                  render={({ field }) => (
                    <TextField
                      variant="outlined"
                      margin="normal"
                      required
                      fullWidth
                      name="Wisdom"
                      label="Wisdom"
                      id="wisdom"
                      type="number"
                      {...field}
                    />
                  )}
                />
              </Grid>
              <Grid item xs={6}>
                <Controller
                  name="Charisma"
                  control={control}
                  render={({ field }) => (
                    <TextField
                      variant="outlined"
                      margin="normal"
                      required
                      fullWidth
                      name="Charisma"
                      label="Charisma"
                      id="charisma"
                      type="number"
                      {...field}
                    />
                  )}
                />
              </Grid>
            </Grid>

            {/* Submit */}
            <Button
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
              className={classes.submit}
            >
              Create
            </Button>
          </form>
        </div>
      ) : (
        <AlertDialog
          open={open}
          setOpen={setOpen}
          title="Character Created!"
          message="Your character has been successfully created. Enjoy!"
        />
      )}
      <Box mt={8}>
        <Copyright />
      </Box>
    </Container>
  );
}
