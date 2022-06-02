import { Category } from "./category";
import { User } from "./user-interface";

export interface VideosInterface {
    _id: string;
    nombre: string;
    categoria:Category;
    usuario:User;
    url:string;
}
