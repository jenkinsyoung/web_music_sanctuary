export const getVisibleProduct =(currentForm : any, products: any) =>{
    if (currentForm == '') return products
    return products.filter((product:any)=>(
        currentForm.includes(product.form && product.pickup_configuration)
    ))
}

export const getVisibleProductByForm =(currentForm : any, products: any) =>{
    if (currentForm == '') return products
    return products.filter((product:any)=>(
        currentForm.includes(product.form)
    ))
}

export const getVisibleProductByPick =(currentForm : any, products: any) =>{
    if (currentForm == '') return products
    return products.filter((product:any)=>(
        currentForm.includes(product.pickup_configuration)
    ))
}