import { useRouteError } from "react-router-dom";
import styled from "styled-components";

export default function ErrorPage() {
    
    const error = useRouteError();

    return (
        <Container>
            <p className="text">Something bad happend!</p>
            <p>{error.statusText || error.message}</p>
        </Container>
    )
}   

const Container = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;

    .text {
        margin: 0;
        font-size: 40px;
        font-weight: 500;
        width: 500px;
    }
`