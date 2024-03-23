export const getCurrentDateAndYear = () => {
    const date = new Date();
    const day = date.getDate();
    const year = date.getFullYear();
    
    return { day, year };
}

export const getMonthName =()=> {
    const date = new Date();
    const month = date.getMonth();
    date.setMonth(month); 
    return date.toLocaleString('en-US', { month: 'long' });
}

export const getDayName =() => {
    const date = new Date();
    return date.toLocaleString('en-US', { weekday: 'long' });
}