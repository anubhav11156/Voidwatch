import React from 'react'
import { Link } from 'react-router-dom'
import styled from 'styled-components'

function Home() {
  return (
    <Container>
      <p>Find a movie to watch tonight!</p>
      <hr />
      <div style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        marginTop: '20px'
      }}>
        <Link to="/movies">
          <img src="/images/movie.png" width={300} height={300} />
        </Link>
      </div>
    </Container>
  )
}

export default Home

const Container = styled.div`
    margin: 0;
    font-weight: 300;
    font-size: 30px;
    width: 100%;

    p {
      font-size: 32px;
      font-weight: 600;
      margin:0;
    }
`