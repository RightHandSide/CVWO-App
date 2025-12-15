import Header from "./Header.tsx"
import Footer from "./Footer.tsx"
import Content from "./Content.tsx"
import Card from "./Card.tsx"
import Button from "./button/Button.tsx"

function App() {
  return (
    <>
      <Header/>
      <Content/>
      <Button/>
      <Card/> <Card/> <Card/>
      <Footer/>
    </>
  );
}

export default App;