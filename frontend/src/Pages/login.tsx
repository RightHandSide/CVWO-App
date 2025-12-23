import Form from "../Components/Form.tsx";

function Login() {
    return (
        <Form type="Login" fetch="http://localhost:8000/login" to="/threads" />
    );
}

export default Login;