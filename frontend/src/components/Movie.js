import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import styled from 'styled-components'


function Movie() {

    let { id } = useParams()
    const [movie, setMovie] = useState({})

    useEffect(() => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");
        const requestOptions = {
            method: "GET",
            headers: headers,
        }

        fetch(`http://localhost:8080/getOneMovie/${id}`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                console.log(data)
                setMovie(data)
            })
            .catch((error) => {
                console.log(error)
            })
    }, [id])

    if (movie.genres) {
        movie.genres = Object.values(movie.genres)
    } else {
        movie.genres = [] // empty array if it doesn't exists
    }

    return (
        <Container>
            <p className='movieName'>Movie: {movie.title}</p>
            <small><pre><em>Release Date: {movie.release_date}   Runtime: {movie.runtime}   Rated: {movie.mpaa_rating}</em></pre></small>
            <small>{movie.genres.map((g) => {
                return (
                    <span key={g._id} className='badge bg-secondary me-2' style={{
                        color: 'white',
                        marginRight: '10px',
                        fontSize: '11px',
                        fontWeight: '600'
                    }}>{g.genre}</span>
                )
            })}</small>
            <hr />
            {movie.image !== "" &&
                <div>
                    <img src={`https://image.tmdb.org/t/p/w200${movie.image}`} alt='poster' />
                </div>
            }

            <hr />
            <p className='description'>{movie.description}</p>
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

    small {
        font-size: 14px;
    }
`