import React, { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import styled from 'styled-components'

function Movies() {

    const [movies, setMovies] = useState([])

    useEffect(() => {

        const headers = new Headers();
        headers.append("Content-Type", "application/json");
        const requestOptions = {
            method: "GET",
            headers: headers,
        }

        fetch(`http://localhost:8080/getAllMovies`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setMovies(data)
            })
            .catch((error) => {
                console.log(error)
            })
    },[])


    return (
        <Container>
            <p>All the movies!</p>
            <hr />
            <table className='table table-striped table-bordered table-hover'>
                <thead>
                    <tr>
                        <th>Movie</th>
                        <th>Release Date</th>
                        <th>Rating</th>
                    </tr>
                </thead>
                <tbody>
                    {movies.map((m) => (
                        <tr key={m._id}>
                            <td>
                                <Link to={`/movies/${m._id}`}>{m.title}</Link>
                            </td>
                            <td>
                                {m.release_date}
                            </td>
                            <td>
                                {m.mpaa_rating}
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </Container>
    )
}

export default Movies

const Container = styled.div`
    font-weight: 200;
    width: 99%;
    p {
        font-size: 32px;
        font-weight: 600;
        margin:0;
    }



    thead {
        font-size: 16px;
        font-weight: 300;
        font-family: poppins;
    }

    tbody {
        font-weight: 300;
        cursor: pointer;
    }
`
