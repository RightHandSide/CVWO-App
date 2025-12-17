import pic from "../assets/NGNL.jpg"

function Card() {
    return (
        <div className="card">
            <img className="card-image" src={pic} alt="Profile Picture"></img>
            <h2 className="card-title">LHS</h2>
            <p className="card-text">1 + 1 = 2</p>
        </div>
    );
}

export default Card;