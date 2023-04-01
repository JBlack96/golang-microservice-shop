import {useParams} from "react-router-dom";

export const ItemPage = () => {
    let { id } = useParams();

    return <div>Item Page: ID {id}</div>;
}

export const ItemPageError = () => <div>Item not found!</div>;
