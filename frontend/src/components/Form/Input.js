import React, { forwardRef } from 'react'
import styled from 'styled-components'


const Input = forwardRef((props, ref) => {
    return (
        <Container>
            {/* <label className="form-label" >{props.title}</label> */}
            <input
                type={props.type}
                ref={ref}
                className={props.className}
                id={props.name}
                placeholder={props.placeholder}
                name={props.name}
                onChange={props.onChange}
                autoComplete={props.autoComplete}
                value={props.value}
            />
        </Container>
    )
})

export default Input

const Container = styled.div`
    input {
        border-radius: 5px;
        border: 1px solid gray;
        outline: none;
        height: 2rem;
        padding-left: 10px;
        margin-bottom: 1rem;
    }
`