import React from "react";

interface ItemCardProps {
    id: string
    name: string
    price: number
    img: string
}

export const ItemCard = ({id, name, price, img}: ItemCardProps) =>
    <div key={id} className="card basis-1/3">
        <figure className="px-10 pt-10">
            <img src={img} alt={name} className="rounded-xl"/>
        </figure>
        <div className="card-body items-center text-center">
            <h2 className="card-title">{name}</h2>
            <p>Price: £{price}</p>
            <div className="card-actions">
                <button className="btn btn-primary">View</button>
            </div>
        </div>
    </div>;
