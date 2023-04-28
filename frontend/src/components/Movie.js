import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import styled from 'styled-components' 


function Movie() {

    let {id} = useParams()
    const [movies, setMovies] = useState({})

    useEffect(() => {
        let movie = {
            id: 1,
            title: "Splinter Cell",
            relase_date: "2002-02-11",
            runtime: 112,
            mppa_rating: "R",
            description: "Some description"
        }
        setMovies(movie)
    },[id])

  return (
    <Container>
        <p className='movieName'>Movie: {movies.title}</p>
        <small><em>Release Date: {movies.relase_date} Runtime: {movies.runtime} Rated:{movies.mppa_rating}</em></small>
        <hr />
        <p className='description'>{movies.description}</p>
    </Container>
  )
}

export default Movie

const Container = styled.div`
    .movieName {
        font-size: 32px;
        font-weight: 600;
        margin:0;
    }

    .descrtiption {
        font-size: 18px;
    }
`