class SessionsController < ApplicationController
  
  def show
    @user = User.new
  end

end
