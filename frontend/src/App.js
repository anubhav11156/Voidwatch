import styled from 'styled-components';
import './App.css';
import { Link, Outlet, useNavigate } from 'react-router-dom';
import { useState } from 'react';
import { useEffect } from 'react';
import { useCallback } from 'react';

function App() {

  const [jwtToken, setJwtToken] = useState("")
  // empty -> user not logged in else logged in

  const navigate = useNavigate();

  const [tickInterval, setTickInterval] = useState()




  const logOutHandler = () => {
    const requestOptions = {
      method: "GET",
      credentials: "include"
    }

    fetch(`/logout`, requestOptions)
      .catch(error => {
        console.log("this")
        console.log("error logging out", error)
      })
      .finally(() => {
        console.log('logged out')
        setJwtToken("")
        toggleRefresh(false)
      })
    navigate("/login")
  }

  const toggleRefresh = useCallback((status) => {
    // if not ticking
    if (status) {
      let i = setInterval(() => {
        const requestOptions = {
          method: "GET",
          credentials: "include"
        }

        fetch(`/refresh`, requestOptions)
          .then((response) => response.json())
          .then((data) => {
            if (data.access_token) {
              setJwtToken(data.access_token)
            }
          })
          .catch(error => {
            console.log("User is not logged in! ")
          })
      }, 600000); // 10 min
      setTickInterval(i)
    } else {
      setInterval(null)
      clearInterval(tickInterval)
    }
  }, [tickInterval])

  useEffect(() => {
    if (jwtToken === "") {
      // get a refresh token
      const requestOptions = {
        method: "GET",
        credentials: "include"
      }

      fetch(`/refresh`, requestOptions)
        .then((response) => response.json())
        .then((data) => {
          if (data.access_token) {
            setJwtToken(data.access_token)
            toggleRefresh(true)
          }
        })
        .catch(error => {
          console.log("User is not logged in! ")
        })
    }
  }, [jwtToken, toggleRefresh])



  return (
    <div className="App">
      <div className='heading'>
        <p className='text'>SubWatch</p>
        {jwtToken === ""
          ? <Link to="/login" className='login-btn'>Login</Link>
          : <a href='#!' className='logut-btn' onClick={logOutHandler}>Logout</a>
        }
      </div>
      <hr />
      <Below>
        <div className='left'>
          <Menu>
            <Link to="/" className='menu-list'>Home</Link>
            <Link to="/movies" className='menu-list'>Movies</Link>
            <Link to="/genres" className='menu-list'>Genres</Link>
            {jwtToken !== "" &&
              <>
                <Link to="/admin/movie/0" className='menu-list'>Add Movie</Link>
                <Link to="/manage-catalouge" className='menu-list'>Manqge Catalogue</Link>
                <Link to="/graphql" className='menu-list' style={{
                  border: "none"
                }}>Graph QL</Link>
              </>
            }
          </Menu>
        </div>
        <div className='right'>
          <Content>
            <Outlet context={{
              jwtToken, setJwtToken, toggleRefresh
            }} />
          </Content>
        </div>
      </Below>
    </div >
  );
}

export default App;

const Below = styled.div`
  height: 100%;
  width: 100%;
  display: flex;

  .left {
    width:15rem;
    height: 100%;
  }

  .right {
    flex: 1;
  }
`

const Menu = styled.div`

  margin-left: 1rem;
  /* height: 15rem; */
  width: 12rem;
  border: 1px solid gray;
  display: flex;
  flex-direction: column;
  border-radius: 10px;
  overflow: hidden;
  
  .menu-list {
    /* flex: 1; */
    height: 40px;
    border-bottom: 1px solid gray;
    width: 100%;
    display: flex;
    justify-content: start;
    align-items: center;
    cursor: pointer;
    text-decoration: none;
    color: black;
    padding-left: 20px;

    &:hover {
      background-color: lightgray;
    }
  }
`
const Content = styled.div`
  flex: 1;
  margin-left: 20px;
`


