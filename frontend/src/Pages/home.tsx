import { Link } from "react-router-dom";

function Home() {
    return (
        <>
            <h1>Home</h1>
            <p>Welcome to Home</p>
            <Link to="/login">Login</Link>
            <hr/>
            <Link to="/register">Register</Link>
        </>
    );
}

export default Home;