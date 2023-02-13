Rails.application.routes.draw do
  resources :quotes
  resources :users
  resource :dashboard

  root "welcome#index"
end
