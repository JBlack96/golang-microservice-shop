import React from 'react';

function App() {
    const items = [
        {id: '99999', name: 'Toothbrush', price: 9.99, imgSrc: 'https://s.alicdn.com/@sc04/kf/He6ab2b71f3304b44a49ad19575a004389.jpg_280x280.jpg',},
        {id: '999919', name: 'Toothbrush', price: 9.99, imgSrc: 'https://s.alicdn.com/@sc04/kf/He6ab2b71f3304b44a49ad19575a004389.jpg_280x280.jpg',},
        {id: '999991', name: 'Toothbrush', price: 9.99, imgSrc: 'https://s.alicdn.com/@sc04/kf/He6ab2b71f3304b44a49ad19575a004389.jpg_280x280.jpg',}
    ];

    return (
        <div>
            <div className="navbar bg-base-100">
                <div className="navbar-start"></div>
                <div className="navbar-center">
                    <a className="btn btn-ghost normal-case text-xl">Golang-Microservice-Shop</a>
                </div>
                <div className="navbar-end"></div>
            </div>
            <div className="flex flex-wrap m-auto max-w-screen-2xl">
                {items.map(item =>
                    <div key={item.id} className="card basis-1/3">
                        <figure className="px-10 pt-10">
                            <img src={item.imgSrc} alt="Shoes" className="rounded-xl"/>
                        </figure>
                        <div className="card-body items-center text-center">
                            <h2 className="card-title">{item.name}</h2>
                            <p>Price: £{item.price}</p>
                            <div className="card-actions">
                                <button className="btn btn-primary">View</button>
                            </div>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
}

export default App;
