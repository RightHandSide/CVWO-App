function Header() {
    return (
        <header>
            <h1>My Site</h1>
            <nav>
                <ul>
                    <li><a href="a">Home</a></li>
                    <li><a href="b">About</a></li>
                    <li><a href="c">Services</a></li>
                    <li><a href="d">Content</a></li>
                    {/* Hyperlink to ./href */}
                </ul>
            </nav>
            <hr></hr>
        </header>
    );
}

export default Header;