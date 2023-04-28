import React, { useState } from 'react'
import styled from 'styled-components'
import Input from './Form/Input'
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { useNavigate, useOutletContext } from 'react-router-dom'

function Login() {

    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")

    const { setJwtToken } = useOutletContext();
    const { toggleRefresh } = useOutletContext();

    const navigate = useNavigate();

    const submitHandler = (event) => {
        event.preventDefault();

        // authenticate user against backend
        let payload = {
            email: email,
            password: password
        }
        const requestOptions = {
            method: "POST",
            Headers: {
                'Content-Type': 'application/json'
            },
            credentials: 'include',
            body: JSON.stringify(payload),
        }

        fetch(`/authenticate`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                // check for error first
                if (data.error) {
                    console.log('error is : ', data.error)
                    console.log('msg: ', data.message)
                } else {
                    setJwtToken(data.access_token)
                    console.log(data.access_token)
                    toggleRefresh(true)
                    navigate('/')
                }
            })
            .catch(error => {
                console.log(error)
            })

    }

    return (
        <Container>
            <div className='login-heading'>
                <p>Login</p>
                <hr />
            </div>
            <LoginForm>
                <div className='form-container'>
                    <form>
                        <Input
                            placeholder={`Username`}
                            type={`email`}
                            name="email"
                            autoComplete="email-new"
                            onChange={(event) => setEmail(event.target.value)}
                        />
                        <Input
                            placeholder={`Password`}
                            type={`password`}
                            name="email"
                            autoComplete="password-new"
                            onChange={(event) => setPassword(event.target.value)}
                        />
                        <button className='submitBtn' onClick={submitHandler}>Submit</button>
                    </form>
                </div>
            </LoginForm>
            <ToastContainer
            // autoClose={1000}
            // hideProgressBar={true}
            />
        </Container>
    )
}

export default Login

const Container = styled.div`
    display: flex;
    flex-direction: column;
    
    .login-heading {
    font-weight: 200;
    width: 99%;

    p {
        font-size: 32px;
        font-weight: 600;
        margin:0;
    }
    }

`

const LoginForm = styled.div`
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
    color: black;

    .form-container {
        margin-top: 3rem;
        padding-top: 2rem;
        height: 25rem;
        width: 25rem;
        border: 1px solid gray;
        border-radius: 10px;
        margin-right: 23rem;
        display: flex;
        justify-content: center;
        align-items: start;

        .submitBtn {
            border: none;
            background-color: blue;
            color: white;
            cursor: pointer;
            width: 5rem;
            height: 2rem;
            border-radius: 5px;
            box-shadow: 0px 1px  10px gray;
            transition: opacity 0.15s ;

            &:hover {
                opacity: 0.8;
            }
        }
    }
`