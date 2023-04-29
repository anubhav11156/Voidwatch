import React from 'react'
import { useEffect } from 'react'
import { useState } from 'react'
import { useNavigate, useOutletContext, useParams } from 'react-router-dom'
import styled from 'styled-components'
import Input from './Form/Input'
import Select from './Form/Select'
import Textarea from './Form/Textarea'

function EditMovie() {

  const navigate = useNavigate()
  const { jwtToken } = useOutletContext()

  const [error, setError] = useState(null)
  const [errors, setErrors] = useState([])
  const [movie, setMovie] = useState({
    id: 0,
    title: "",
    release_data: "",
    runtime: "",
    mpaa_rating: "",
    description: ""
  })

  const mpaaOptions = [
    {id:"G", value:"G"},
    {id:"PG", value:"PG"},
    {id:"PG13", value:"PG13"},
    {id:"R", value:"R"},
    {id:"18A", value:"18A"},
    {id:"NC17", value:"NC17"},
  ]

  let { id } = useParams();

  useEffect(() => {
    if (jwtToken === "") {
      navigate("/login")
      return
    }
  }, [jwtToken, navigate])

  const handleSubmit = (event) => {
    event.preventDefault()
  }

  // using only one fucntion to change all the propertise of movie object
  const handleChange = () => (event) => {
    let value = event.target.value;
    let name = event.target.name

    setMovie({
      ...movie,
      [name]: value
    })
  }

  return (
    <Container>
      <div className='edit-heading'>
        <p>Add/Edit Movie</p>
        <hr />
      </div>

      <div className='form-container'>
        <pre>{JSON.stringify(movie, null, 3)}</pre>

        <form onSubmit={handleSubmit}>
          <input type='hidden' name='id' value={movie.id} />

          <Input
            title={"Title"}
            className={"form-control"}
            type={"title"}
            name={"title"}
            value={movie.title} // basically binding
            onChange={handleChange("title")}
            placeholder={"Movie Title"}
          />
          <Input
            title={"Release Date"}
            className={"form-control"}
            type={"date"}
            name={"release_data"}
            value={movie.release_data}
            onChange={handleChange("release_date")}
          />
          <Input
            title={"Runtime"}
            className={"form-control"}
            type={"text"}
            name={"runtime"}
            value={movie.runtime}
            onChange={handleChange("runtime")}
            placeholder={"Runtime"}
          />
          <Select
            name={"mpaa_rating"}
            option={mpaaOptions}
            onChange={handleChange("mpaa_rating")}
            placeHolder={"Rating"}
          />
          <Textarea 
            title="Description"
            name={"description"}
            value={movie.description}
            rows="3"
            onChange={handleChange("description")}
          />
        </form>
      </div>


    </Container>
  )
}

export default EditMovie

const Container = styled.div`
  width: 99%;
  display: flex;
  flex-direction: column;
  justify-content: start;
  align-items: start;
    p {
       font-size: 32px;
       font-weight: 600; 
       margin:0;
   }

  .edit-heading {
    width: 99%;
    height: 5rem;
  } 

  .form-container {
    width: 28rem;
    flex: 1;

    Select {
      width: 100%;
      cursor: pointer;
    }
  }

`
