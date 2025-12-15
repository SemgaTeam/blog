import { Admin, Resource } from "react-admin";
import { Layout } from "./Layout";

import dataProvider from "./dataProvider.ts";
import authProvider from "./authProvider.ts";

import PostList from "./components/posts/PostList";
import PostShow from "./components/posts/PostShow";
import PostEdit from "./components/posts/PostEdit";

const App = () => (
  <Admin
    layout={Layout}
    dataProvider={dataProvider}
    authProvider={authProvider}
    disableTelemetry
  >
    <Resource name="post" list={PostList} show={PostShow} edit={PostEdit} />
    <Resource name="user" />
  </Admin>
);

export default App;
