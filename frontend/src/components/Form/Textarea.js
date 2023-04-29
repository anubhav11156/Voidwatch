const Textarea = (props) => {
    return (
        <div className="mb-3">
            <label htmlFor={props.name} className="form-label">
                {props.title}
            </label>
            <textarea
                className="form-control"
                id={props.name}
                name={props.name}
                value={props.value}
                onChange={props.onChange}
                rows={props.row}
            />
            <div className={props.erroDiv}>{props.erroMsg}</div>
        </div>
    )
}

export default Textarea