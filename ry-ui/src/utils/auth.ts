import Cookies from "js-cookie";

const TokenKey = "Admin-Token";

export const getToken = () => {
    return Cookies.get(TokenKey);
};

export const setToken = (token: string) => {
    return Cookies.set(TokenKey, token);
};

export const removeToken = () => {
    Cookies.remove("Authorization");
    return Cookies.remove(TokenKey);
};
