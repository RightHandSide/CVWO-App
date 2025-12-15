import styles from "./Button.module.css"
import pic from "../assets/NGNL.jpg"

function Button() {
    return (
        <>
        <button className={styles.button} onClick={() => console.log("Clicked")}>
            <img className={styles["button-image"]} src={pic}></img>
        </button>
        <hr></hr>
        </>
    );
}

export default Button;