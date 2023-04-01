import React from 'react';

import {Navbar} from "../Components/Navbar";
import {ItemCard} from "../Components/ItemCard";

function Home() {
    const items = [
        {id: '99999', name: 'Toothbrush', price: 9.99, imgSrc: 'https://s.alicdn.com/@sc04/kf/He6ab2b71f3304b44a49ad19575a004389.jpg_280x280.jpg',},
        {id: '999919', name: 'Toothbrush', price: 9.99, imgSrc: 'https://s.alicdn.com/@sc04/kf/He6ab2b71f3304b44a49ad19575a004389.jpg_280x280.jpg',},
        {id: '999991', name: 'Toothbrush', price: 9.99, imgSrc: 'https://s.alicdn.com/@sc04/kf/He6ab2b71f3304b44a49ad19575a004389.jpg_280x280.jpg',}
    ];

    return (
        <div>
            <Navbar />
            <div className="flex flex-wrap m-auto max-w-screen-2xl">
                {items.map(item => <ItemCard id={item.id} name={item.name} price={item.price} img={item.imgSrc}/>)}
            </div>
        </div>
    );
}

export default Home;
