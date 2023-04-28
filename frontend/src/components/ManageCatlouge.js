import {React, useState, useEffect} from 'react'
import styled from 'styled-components'
import { Link, useNavigate, useOutletContext } from 'react-router-dom'

function ManageCatlouge() {
  const [movies, setMovies] = useState([])
  const {jwtToken} = useOutletContext()
  
  const navigate = useNavigate();

  useEffect(() => {
      if (jwtToken === ""){
        navigate("/login")
        return
      }

      const headers = new Headers();
      headers.append("Content-Type", "application/json");
      headers.append("Authorization", "Bearer " + jwtToken)

      const requestOptions = {
          method: "GET",
          headers: headers,
      }

      fetch(`http://localhost:8080/admin/movies`, requestOptions)
          .then((response) => response.json())
          .then((data) => {
              console.log('data : ', data)
              setMovies(data)
          })
          .catch((error) => {
              console.log(error)
          })
  },[jwtToken, navigate])

  return (
      <Container>
          <p>Manage Catlouge</p>
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
                      <tr key={m.id}>
                          <td>
                              <Link to={`/admin/movies/${m.id}`}>{m.title}</Link>
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

export default ManageCatlouge

const Container = styled.div`
   p {
        font-size: 32px;
        font-weight: 600;
        margin:0;
    }
`