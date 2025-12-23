import { useNavigate } from "react-router-dom";
import { Box, Button, Stack, Typography } from "@mui/material";

function Home() {
    const naviagte = useNavigate();
    
    return (
        <Box
            sx={{
                minHeight: "100vh",
                display: "flex",
                justifyContent: "center",
                alignItems: "center"
            }}>
                <Box textAlign="center">
                    <Typography variant="h3" gutterBottom>
                        Home
                    </Typography>
                    <Stack direction="row" spacing={2} justifyContent="center">
                        <Button variant="contained" onClick={() => naviagte("/login")}>
                            Login
                        </Button>
                        <Button variant="contained" onClick={() => naviagte("/register")}>
                            Register
                        </Button>
                    </Stack>
                </Box>
            </Box>
    );
}

export default Home;