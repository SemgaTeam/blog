import { Admin, Resource, ShowGuesser, EditGuesser } from "react-admin";
import { Layout } from "./Layout";
import dataProvider from "./dataProvider.ts";
import PostList from "./components/posts/PostList";

const App = () => (
  <Admin layout={Layout} dataProvider={dataProvider}>
    <Resource
      name="post"
      list={PostList}
      show={ShowGuesser}
      edit={EditGuesser}
    />
    <Resource name="user" />
  </Admin>
);

export default App;
