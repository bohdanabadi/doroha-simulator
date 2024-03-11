import {Container, Nav, Navbar} from "react-bootstrap";
function NavBarComponent() {
    return (
            <Navbar expand="lg" className="bg-body-tertiary">
                <Container>
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="me-auto justify-content-center w-100">
                            <Nav.Link className="mx-2" href="/">Home</Nav.Link>
                            <Nav.Link className="mx-2" href="/metrics">Metrics</Nav.Link>
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
    );
}

export default NavBarComponent