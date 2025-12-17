import Header from "../Component/Header.tsx"
import Footer from "../Component/Footer.tsx"
import Content from "../Component/Content.tsx"
import Card from "../Component/Card.tsx"

function Test() {
  return (
    <>
      <Header/>
      <Content/>
      <Card/> <Card/> <Card/>
      <Footer/>
    </>
  );
}

export default Test;