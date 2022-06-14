import { User } from "./user-interface";

export interface Category {
    categorias:Array<Category>;
    _id:string;
    nombre:string;
    total:number;
    usuario:User;
}
