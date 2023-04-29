const Checkbox = (props) => {
    return (
        <div className="form-check">
            <input
                id={props.name}
                className="form-check-input"
                value={props.value}
                name={props.name}
                onChange={props.onChange}
                checked={props.checked}
            />
            <label className="form-check-label" htmlFor={props.name}>
                {props.title}
            </label>
        </div>
    )
}

export default Checkbox