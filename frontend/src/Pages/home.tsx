import { Link } from "react-router-dom";

function Home() {
    return (
        <div className="center-box">
            <h1>Home</h1>
            <div style={{display: "flex", gap: "8px"}}>
                <Link to="/login"><button>Login</button></Link>
                <Link to="/register"><button>Register</button></Link>
            </div>
        </div>
    );
}

export default Home;