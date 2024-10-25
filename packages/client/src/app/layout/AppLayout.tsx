import { Grid2 } from "@mui/material";
import { AppFooter, AppHeader, AppMain } from "app/components";

export const AppLayout = () => {
  return (
    <Grid2
      container
      display="grid"
      gridTemplateRows="60px 1fr 60px"
      height="100vh"
      width="100wv"
    >
      <AppHeader />
      <AppMain />
      <AppFooter />
    </Grid2>
  );
};
