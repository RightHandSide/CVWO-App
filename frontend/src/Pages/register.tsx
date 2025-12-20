import Form from "../Component/Form.tsx";

function Register() {
    return (
        <Form type="Register" fetch="http://localhost:8000/register" to="/register/completed" />
    );
}

export default Register;